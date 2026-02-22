import Link from "@docusaurus/Link";
import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import Heading from "@theme/Heading";
import clsx from "clsx";
import { useEffect, useState } from "react";
import styles from "./styles.module.css";

const lines = [
  "$ go install github.com/abhishekkushwahaa/go-microx/cmd/go-microx@latest",
  "$ go-microx new",
  "? Select a project template: E-commerce",
  "? Enter project name: myshop",
  "? Select an HTTP Router: Gin",
  "? Select a database: PostgreSQL",
  "? Select an authentication method: JWT",
  "Creating project: myshop...",
  "Project myshop has been successfully created! 🎉",
  "$ cd myshop && docker-compose up --build",
];

export default function CodePreview({ isHero }) {
  const { siteConfig } = useDocusaurusContext();
  const [currentLine, setCurrentLine] = useState(0);
  const [displayText, setDisplayText] = useState([]);

  useEffect(() => {
    if (currentLine < lines.length) {
      const timeout = setTimeout(() => {
        setDisplayText((prev) => [...prev, lines[currentLine]]);
        setCurrentLine((prev) => prev + 1);
      }, 1000);
      return () => clearTimeout(timeout);
    }
  }, [currentLine]);

  return (
    <section className={clsx(styles.codeSection, isHero && styles.heroMode)}>
      {isHero && (
        <div className={styles.heroBackground}>
          <div className={styles.heroCircle1}></div>
          <div className={styles.heroCircle2}></div>
        </div>
      )}
      <div className="container">
        {isHero && (
          <div className={styles.heroHeader}>
            <Heading as="h1" className={styles.heroTitle}>
              {siteConfig.title}
            </Heading>
            <p className={styles.heroSubtitle}>{siteConfig.tagline}</p>
            <div className={styles.heroButtons}>
              <Link
                className={clsx("button button--lg", styles.buttonPrimary)}
                to="/docs/intro"
              >
                Get Started
              </Link>
              <Link
                className={clsx("button button--lg", styles.buttonOutline)}
                href="https://github.com/abhishekkushwahaa/go-microx"
              >
                GitHub
              </Link>
            </div>
          </div>
        )}
        <div className={styles.terminalContainer}>
          <div className={styles.terminal}>
            <div className={styles.terminalHeader}>
              <div className={styles.buttons}>
                <span className={styles.dotRed}></span>
                <span className={styles.dotYellow}></span>
                <span className={styles.dotGreen}></span>
              </div>
              <div className={styles.terminalTitle}>bash — go-microx</div>
            </div>
            <div className={styles.terminalBody}>
              {displayText.map((line, i) => (
                <div key={i} className={styles.terminalLine}>
                  <span className={line.startsWith("$") ? styles.prompt : ""}>
                    {line}
                  </span>
                </div>
              ))}
              {currentLine < lines.length && (
                <span className={styles.cursor}>_</span>
              )}
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
