import { Download, Play, Terminal } from "lucide-react";
import styles from "./styles.module.css";

const steps = [
  {
    title: "1. Install CLI",
    description: "Install the go-microx CLI tool globally using Go install command.",
    icon: Download,
  },
  {
    title: "2. Create Project",
    description: "Run 'go-microx new' and follow the interactive wizard to scaffold your app.",
    icon: Terminal,
  },
  {
    title: "3. Run & Build",
    description: "Navigate to your project folder and run with Docker or natively.",
    icon: Play,
  },
];

export default function HowItWorks() {
  return (
    <section className={styles.section}>
      <div className="container">
        <h2 className={styles.title}>Simple Workflow</h2>
        <div className={styles.stepsGrid}>
          {steps.map((step, i) => (
            <div key={i} className={styles.stepCard}>
              <div className={styles.iconWrapper}>
                <step.icon size={32} />
              </div>
              <h3 className={styles.stepTitle}>{step.title}</h3>
              <p className={styles.stepDescription}>{step.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
