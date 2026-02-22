import { Code2, Cpu, Layers, Rocket, ShieldCheck, Zap } from "lucide-react";
import "./HomepageFeatures.css";

const features = [
  {
    title: "Industry Templates",
    icon: Layers,
    description:
      "Ready-to-use templates for E-commerce, Video Streaming, and Food Delivery architectures.",
  },
  {
    title: "Interactive CLI",
    icon: Rocket,
    description: "Launch your project with an intuitive wizard. No more hunting through documentation for flags.",
  },
  {
    title: "Database Freedom",
    icon: Cpu,
    description:
      "Seamless integration with PostgreSQL, MySQL, MongoDB, and SQLite out of the box.",
  },
  {
    title: "Router Selection",
    icon: Zap,
    description:
      "Choose from Gin, Fiber, Echo, or Chi. Use the framework that fits your team's expertise.",
  },
  {
    title: "Production Ready",
    icon: ShieldCheck,
    description: "Generates clean, modular code with Docker and Docker Compose support for easy deployment.",
  },
  {
    title: "Open Source",
    icon: Code2,
    description: "Community driven and open for contributions. Extend it to fit your unique microservices needs.",
  },
];

export default function HomepageFeatures() {
  return (
    <section className="features-section">
      <h2 className="features-title">
        Build Scalable Architectures with{" "}
        <span className="highlight">go-microx</span>
      </h2>

      <div className="features-grid">
        {features.map((feature, idx) => (
          <div key={idx} className="feature-card">
            <div className="feature-icon">
              <feature.icon size={40} />
            </div>
            <h3 className="feature-title">{feature.title}</h3>
            <p className="feature-description">{feature.description}</p>
          </div>
        ))}
      </div>
    </section>
  );
}
