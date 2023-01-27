import styles from '@/components/atoms/Link.module.css';

type LinkProps = {
  href: string;
  children: React.ReactNode;
  newTab?: boolean;
  className?: string;
  style?: React.CSSProperties;
};

export default function Link({
  href,
  children,
  newTab,
  className = ``,
  style = {},
}: LinkProps) {
  return (
    <a
      href={href}
      className={`${styles.link} ${className}`}
      style={style}
      target={newTab ? `_blank` : `_self`}
      rel={newTab ? `noopener noreferrer` : ``}
    >
      {children}
    </a>
  );
}
