import { Link } from 'react-router-dom';
import AllSpaces from './AllSpaces';
import MySpaceMemberships from './MySpaceMemberships';

export default function() {
  return (
    <div className="flex flex-col p-10">
      <section className="pb-10">
        <AllSpaces />
      </section>
    </div>
  );
}
