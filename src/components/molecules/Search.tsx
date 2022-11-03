import Input from '@/components/atoms/Input';

export default function Search() {
  return (
    <div
      style={{
        display: `flex`,
        alignItems: `center`,
        position: `absolute`,
        left: `50%`,
        width: `100%`,
        transform: `translateX(-50%)`,
        justifyContent: `center`,
      }}
    >
      <Input />
    </div>
  );
}
