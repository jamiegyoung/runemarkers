import SectionHeader from '@/components/atoms/SectionHeader';

type SectionHeaderProps = {
  title: string;
  centered?: boolean;
  children?: JSX.Element | JSX.Element[];
};

export default function ListContainerSection({
  children,
  title,
}: SectionHeaderProps) {
  return (
    <div
      style={{
        display: `flex`,
        flexDirection: `column`,
        alignItems: `center`,
        width: `90%`,
        margin: `10px 0 10px 0`,
      }}
    >
      <SectionHeader text={title} />
      {children}
    </div>
  );
}
