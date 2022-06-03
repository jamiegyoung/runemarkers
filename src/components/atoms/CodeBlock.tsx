import SyntaxHighlighter from 'react-syntax-highlighter';
import { tomorrowNight } from 'react-syntax-highlighter/dist/cjs/styles/hljs';
import styles from './CodeBlock.module.css';

export default function CodeBlock({ text }: { text: string }) {
  return (
    <div className={styles.container}>
      <h2 className={styles.header}>Tile Data:</h2>
      <SyntaxHighlighter
        language="JSON"
        style={tomorrowNight}
        wrapLongLines
        wrapLines
        customStyle={{
          borderRadius: 10,
          backgroundColor: `#3b3b37`,
          boxShadow: `0px 0px 10px rgba(0, 0, 0, 0.5)`,
          padding: 20,
          margin: `0 30px`,
        }}
      >
        {text}
      </SyntaxHighlighter>
    </div>
  );
}
