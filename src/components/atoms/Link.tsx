import styles from '@/components/atoms/Link.module.css';

type LinkProps = {
  href: string;
  children: React.ReactNode;
  newTab?: boolean;
};

export default function Link({ href, children, newTab }: LinkProps) {
  return (
    <a
      href={href}
      className={styles.link}
      target={newTab ? `_blank` : `_self`}
      rel={newTab ? `noopener noreferrer` : ``}
    >
      {children}
    </a>
  );
}
