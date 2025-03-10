import clsx from "clsx";
import Heading from "@theme/Heading";
import styles from "./styles.module.css";

const FeatureList = [
  {
    title: "Stop Reinventing, Start Building",
    Svg: require("@site/static/img/micro.svg").default,
    description: (
      <>
        Tired of setting up microservices from scratch every time? Gomicrox
        eliminates the hassle by providing a scalable, production-ready
        boilerplate—so you can focus on what truly matters.
      </>
    ),
  },
  {
    title: "Rapid Development, Zero Overhead",
    Svg: require("@site/static/img/micro.svg").default,
    description: (
      <>
        Gomicrox automates the tedious setup of authentication, database
        connections, and service communication. Spin up a microservices
        architecture in minutes, not hours!
      </>
    ),
  },
  {
    title: "Built for Scale & Performance",
    Svg: require("@site/static/img/micro.svg").default,
    description: (
      <>
        Designed with Golang, Gomicrox ensures high performance, reliability,
        and modularity—perfect for modern cloud-native applications.
      </>
    ),
  },
];

function Feature({ Svg, title, description }) {
  return (
    <div className={clsx("col col--4")}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
