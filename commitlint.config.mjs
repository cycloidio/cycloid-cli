// Cycloid commit convention — feeds Linear releases (changie is being retired).
//
// Conventional Commits + a Linear issue key on release-note-bearing commits.
// Release-note types (feat/fix/perf) MUST reference a Linear issue key (e.g. BE-123)
// in the subject or body, so the release pipeline can link the issue to the Linear
// release. Maintenance types (chore/build/ci/docs/style/test/refactor/revert) are
// exempt — they land in the "Maintenance" bucket and need no key.
//
// ESM (.mjs) is required by wagoid/commitlint-github-action@v6.
// NOTE: intended to be extracted into a shared @cycloid/commitlint-config package
// once the monorepo lands; kept inline per-repo for now.

const RELEASE_NOTE_TYPES = ['feat', 'fix', 'perf'];
const LINEAR_KEY = /\b[A-Z]{2,}-\d+\b/;

export default {
  extends: ['@commitlint/config-conventional'],
  plugins: [
    {
      rules: {
        'linear-key-on-release-types': (parsed) => {
          const { type, header, body } = parsed;
          if (!type || !RELEASE_NOTE_TYPES.includes(type)) return [true];
          const hasKey = LINEAR_KEY.test(`${header || ''}\n${body || ''}`);
          return [
            hasKey,
            `a "${type}" commit must reference a Linear issue key (e.g. BE-123) ` +
              `in the subject or body so it can be linked to the Linear release`,
          ];
        },
      },
    },
  ],
  rules: {
    'linear-key-on-release-types': [2, 'always'],
  },
};
