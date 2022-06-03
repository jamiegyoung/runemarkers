import Input from '@/components/atoms/Input';
// import SearchIcon from "../atoms/SearchIcon";

export default function Search() {
  return (
    <div
      style={{
        display: `flex`,
        alignItems: `center`,
        marginBottom: 20,
      }}
    >
      <Input />
      {/* <SearchIcon /> */}
    </div>
  );
}
