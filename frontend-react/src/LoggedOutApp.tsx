import Button from "./Button";

export default function LoggedOutApp() {
  return (
    <div className="grid h-screen place-items-center">
      <div className="flex flex-col items-center">
        <img alt="octopus" width="300px" src="octopus1.png" />
        <img
          alt="octopus junior text"
          width="300px"
          src="octopus-junior-text.png"
        />

        <a
          href={`/auth/login?from=${encodeURIComponent(window.location.href)}`}
        >
          <Button color="blue">login</Button>
        </a>
      </div>
    </div>
  );
}
