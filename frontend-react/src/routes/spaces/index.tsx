// import { useAllSpacesQuery } from '../../generated-types';
// import { Link } from "react-router-dom";
// import JoinSpaceButton from '../../components/JoinSpaceButton';
import { useContext } from 'react';
import { PersonIdContext } from '../../contexts';
import AllSpaces from './AllSpaces';
import MySpaceMemberships from './MySpaceMemberships';

export default function() {
  const personId = useContext(PersonIdContext)

  return (
    <div className="flex flex-col">
      <section>
        <h1 className="text-xl">my spaces</h1>
        <MySpaceMemberships />
      </section>

      <section>
        <h1 className="text-xl">all spaces</h1>
        <AllSpaces />
      </section>
    </div>
  );
}
