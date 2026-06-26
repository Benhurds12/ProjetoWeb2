import "./style.css";

export default function Navbar({ titulo }) {
  return (
    <header className="navbar">
      <h1>{titulo}</h1>
    </header>
  );
}