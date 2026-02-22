import Link from "@docusaurus/Link";
import styles from "./styles.module.css";

export default function CTA() {
  return (
    <section className={styles.section}>
      <div className="container">
        <div className={styles.card}>
          <h2 className={styles.title}>Ready to Build Your Next Microservice?</h2>
          <p className={styles.subtitle}>
            Join the community and start scaffolding production-ready architectures in seconds.
          </p>
          <div className={styles.buttons}>
            <Link className="button button--primary button--lg" to="/docs/intro">
              Get Started Now
            </Link>
            <Link className="button button--outline button--secondary button--lg" href="https://github.com/abhishekkushwahaa/go-microx">
              Star on GitHub
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
}
