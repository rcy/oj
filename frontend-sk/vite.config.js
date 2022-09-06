import { sveltekit } from '@sveltejs/kit/vite';

const config = {
  plugins: [sveltekit()],
  server: {
    proxy: {
      '/auth': {
        target: 'http://localhost:5000'
      },
      '/graphql': {
        target: 'http://localhost:5000'
      },
      '/graphiql': {
        target: 'http://localhost:5000'
      }
    }
  }
};

export default config;
