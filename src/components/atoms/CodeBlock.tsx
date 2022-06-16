import { useState } from 'react';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { tomorrowNight } from 'react-syntax-highlighter/dist/cjs/styles/hljs';
import styles from './CodeBlock.module.css';

const TRUNCATE_LENGTH = 700;

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
          wordBreak: `break-all`,
        }}
      >
        {text.length > TRUNCATE_LENGTH && !showingAll
          ? `${text.slice(0, TRUNCATE_LENGTH)}...`
          : text}
      </SyntaxHighlighter>
      <code
        style={{
          cursor: `pointer`,
        }}
        onClick={() => {
          setShowingAll(!showingAll);
        }}
      >
        {text.length > TRUNCATE_LENGTH
          ? showingAll
            ? `Show Less`
            : `Show All (${
                text.length - TRUNCATE_LENGTH
              } more characters, this may take a while to
        load)`
          : null}
      </code>
    </div>
  );
}
