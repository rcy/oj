import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "../backend/schema.graphql",
  documents: [
    "src/**/*.graphql",
  ],
  generates: {
    "src/generated-types.ts": {
      plugins: ['typescript', 'typescript-operations', 'typescript-react-apollo']
    },
  }
};

export default config;
