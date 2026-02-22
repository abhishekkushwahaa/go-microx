import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import CodePreview from "@site/src/components/CodePreview";
import CTA from "@site/src/components/CTA";
import HomepageFeatures from "@site/src/components/HomepageFeatures";
import HowItWorks from "@site/src/components/HowItWorks";
import Layout from "@theme/Layout";


export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`${siteConfig.title} | Build Microservices with Ease`}
      description="Gomicrox helps you build scalable microservices effortlessly."
    >
      <main>
        <CodePreview isHero={true} />
        <HomepageFeatures />
        <HowItWorks />
        <CTA />
      </main>
    </Layout>
  );
}
