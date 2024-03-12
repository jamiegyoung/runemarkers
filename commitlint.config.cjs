module.exports = {
  extends: ['@commitlint/config-conventional'],
  // Ignore Dependabot commits
  ignores: [(message) => /^Bumps \[.+]\(.+\) from .+ to .+\.$/m.test(message)],
};
