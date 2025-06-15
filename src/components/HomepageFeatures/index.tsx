import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Simples',
    Svg: require('@site/static/img/undraw_docusaurus_mountain.svg').default,
    description: (
      <>
        Go oferece uma sintaxe limpa e minimalista que elimina a verbosidade 
        comum em outras linguagens. Para um desenvolvedor Java, isso significa
        menos c&oacute;digo para escrever e uma curva de aprendizado incrivelmente
        suave para se tornar produtivo em pouqu&iacute;ssimo tempo.
      </>
    ),
  },
  {
    title: 'Rápido',
    Svg: require('@site/static/img/gopher-flash-sm.svg').default,
    description: (
      <>
        Projetado para performance, Go compila diretamente para c&oacute;digo de m&aacute;quina, 
        resultando em inicialização quase instant&acirc;nea e um desempenho excepcional. Esque&ccedil;a 
        o tempo de aquecimento da JVM; suas aplicações Go rodam com velocidade máxima desde o 
        primeiro momento.
      </>
    ),
  },
  {
    title: 'Poderoso',
    Svg: require('@site/static/img/gopher-strong-sm.svg').default,
    description: (
      <>
        A concorr&ecirc;ncia é uma cidad&atilde; de primeira classe em Go, com 
        goroutines e channels integrados &agrave; linguagem. Construir sistemas complexos 
        e concorrentes se torna uma tarefa trivial, permitindo que voc&ecirc; aproveite ao 
        m&aacute;ximo os processadores multi-core modernos sem esfor&ccedil;o.
      </>
    ),
  },
];

function Feature({title, Svg, description}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
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

export default function HomepageFeatures(): ReactNode {
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
