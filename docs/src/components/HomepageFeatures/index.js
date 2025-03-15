import { Cpu, Rocket, Layers, Code2, ShieldCheck, Zap } from "lucide-react";
import "./HomepageFeatures.css";

const features = [
  {
    title: "Instant Setup",
    icon: Layers,
    description:
      "Spin up microservices instantly with GoMicrox's powerful scaffolding.",
  },
  {
    title: "Blazing Fast Deployment",
    icon: Rocket,
    description: "Deploy production-ready microservices in just a few clicks.",
  },
  {
    title: "Optimized Performance",
    icon: Cpu,
    description:
      "Leverage Golang’s speed for low-latency, high-performance apps.",
  },
  {
    title: "Rock-Solid Security",
    icon: ShieldCheck,
    description:
      "Built-in security ensures your architecture is always protected.",
  },
  {
    title: "Seamless Integrations",
    icon: Code2,
    description: "Integrate with Docker, Kubernetes, and CI/CD out of the box.",
  },
  {
    title: "Lightning Fast Prototyping",
    icon: Zap,
    description:
      "Kickstart projects with pre-built templates and reduce dev time.",
  },
];

export default function HomepageFeatures() {
  return (
    <section className="features-section">
      <h2 className="features-title">
        Supercharge Your Microservices with{" "}
        <span className="highlight">GoMicrox</span>
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
