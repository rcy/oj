import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "../backend/schema.graphql",
  documents: [
    "src/**/*.graphql",
    "src/**/*.tsx",
  ],
  ignoreNoDocuments: true, // for better experience with the watcher
  generates: {
    "src/generated-types.ts": {
      plugins: ['typescript', 'typescript-operations', 'typescript-react-apollo']
    },
    'src/gql/': {
      preset: 'client',
      plugins: []
    },
  },
};

export default config;
