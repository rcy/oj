export default function Debug(x: any) {
  return <pre>{JSON.stringify(x, null, 2)}</pre>;
}
