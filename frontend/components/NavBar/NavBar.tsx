import * as styles from './NavBar.module.css';

export const NavBar: React.FC = () => {
  return (
    <nav className={styles.root}>
      <a href="/legions">Legions</a>
      <a href="/primarchs">Primarchs</a>
      <a href="/battles">Battles</a>
      <a href="/tactics">Tactics</a>
      <a href="/lore">Lore</a>
    </nav>
  );
};
