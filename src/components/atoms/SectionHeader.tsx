type SectionHeaderProps = {
  text: string;
};

export default function SectionHeader({ text }: SectionHeaderProps) {
  return (
    <h2
      style={{
        margin: `10px 10px 20px 10px`,
        alignSelf: `flex-start`,
      }}
    >
      {text}
    </h2>
  );
}
