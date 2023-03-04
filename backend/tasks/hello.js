export async function hello(payload, helpers) {
  const { name } = payload;
  helpers.logger.info(`Hello, ${name}`);
};
