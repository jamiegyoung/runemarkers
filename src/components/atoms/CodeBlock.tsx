import { useState } from 'react';
import SyntaxHighlighter from 'react-syntax-highlighter';
import { tomorrowNight } from 'react-syntax-highlighter/dist/cjs/styles/hljs';
import styles from './CodeBlock.module.css';

const TRUNCATE_LENGTH = 700;

export default function CodeBlock({
  text,
  truncateLength = TRUNCATE_LENGTH,
}: {
  text: string;
  truncateLength?: number;
}) {
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
        {text.length > truncateLength && !showingAll
          ? `${text.slice(0, truncateLength)}...`
          : text}
      </SyntaxHighlighter>
      {text.length > truncateLength ? (
        <code
          id={styles.toggleShowAll}
          onClick={() => {
            setShowingAll(!showingAll);
          }}
        >
          {showingAll
            ? `Show Less`
            : `Show All (${
                text.length - truncateLength
              } more characters, this may take a while to
        load)`}
        </code>
      ) : null}
    </div>
  );
}
