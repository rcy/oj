export default function beep() {
  const snd = new Audio("/chat-alert.mp3");
  snd.play();
}
