#!/usr/bin/env node
/* eslint-disable import/no-commonjs */
/* eslint-disable @typescript-eslint/no-var-requires */
const chalk = require('chalk');
const execa = require('execa');
const micromatch = require('micromatch');

const clientConfig = require('../../../config/clients.config.json');
const GENERATED_FILE_PATTERNS =
  require('../../../config/generation.config').patterns;

const run = async (command, { cwd } = {}) => {
  return (
    (await execa.command(command, { shell: 'bash', all: true, cwd })).all ?? ''
  );
};

function getPatterns() {
  const patterns = GENERATED_FILE_PATTERNS;
  for (const [language, { tests }] of Object.entries(clientConfig)) {
    patterns.push(`tests/output/${language}/${tests.outputFolder}/client/**`);
    patterns.push(`tests/output/${language}/${tests.outputFolder}/methods/**`);
  }
  return patterns;
}

async function preCommit() {
  // when merging, we want to stage all the files
  if ((await run('git merge HEAD')) !== 'Already up to date.') {
    return;
  }

  const stagedFiles = (
    await run('git diff --name-only --cached --diff-filter=d')
  ).split('\n');

  const toUnstage = micromatch.match(stagedFiles, getPatterns());

  for (const file of toUnstage) {
    // eslint-disable-next-line no-console
    console.log(
      chalk.black.bgYellow('[INFO]'),
      `Generated file found, unstaging: ${file}`
    );

    await run(`git restore --staged ${file}`);
  }
}

if (require.main === module && !process.env.CI) {
  preCommit();
}

module.exports = { getPatterns };
