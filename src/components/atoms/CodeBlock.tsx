import { useState } from 'react';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { tomorrowNight } from 'react-syntax-highlighter/dist/cjs/styles/hljs';
import styles from './CodeBlock.module.css';

export default function CodeBlock({ text }: { text: string }) {
  const [showingAll, setShowingAll] = useState(false);
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
        }}
      >
        {text.length > 1000 && !showingAll ? `${text.slice(0, 1000)}...` : text}
      </SyntaxHighlighter>
      <code
        style={{
          cursor: `pointer`,
        }}
        onClick={() => {
          setShowingAll(!showingAll);
        }}
      >
        {text.length > 1000 && !showingAll
          ? `
        Click to view ${text.length - 1000} more characters (may take a while to
        load)`
          : `Click to minimize`}
      </code>
    </div>
  );
}
