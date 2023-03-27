import { Link, Route, Routes, useLocation, useNavigate } from "react-router-dom";
import { VStack, Box, Button, Container, Divider, Heading, Spinner, Text, Flex, HStack, Spacer } from "@chakra-ui/react";
import AdminAddFamilyMember from "../../components/AdminAddFamilyMember";
import { useCurrentUserFamilyQuery } from "../../generated-types";
import PersonCard from "../../components/PersonCard";

export default function OnboardFamily() {
  const location = useLocation()
  const navigate = useNavigate()
  const q = useCurrentUserFamilyQuery()

  async function success(x: any) {
    await q.refetch()
    navigate(`${location}/..`)
  }

  function cancel() {
    navigate(`${location}/..`)
  }

  if (q.loading) {
    return <Spinner />
  }

  console.log(q.data)

  return (
    <Container pb={10}>
      <Heading>My Family</Heading>
      <Divider my="4" />

      <Routes>
        <Route path="/" element={
          <VStack align="left">
            <HStack>
              <Heading size="md" mb="2">Children</Heading>
              <Spacer/>
              <Link to="add-managed">
                <Button size="sm">Create Child Account</Button>
              </Link>
            </HStack>
            <Text>
              With a managed child account, you can monitor your child's activity, control who they interact with on the platform, and set usage limits. This means that you can customize the settings to fit your child's needs and ensure their safety and security.
            </Text>
            <Box>
              {q.data?.currentUser?.family?.familyMemberships.nodes.filter(n => !n.person?.user).map(n => (
                <PersonCard key={n.id} person={n.person} />
              ))}
            </Box>
          </VStack>
        } />
        <Route path="/add-managed" element={
          <AdminAddFamilyMember onSuccess={success} onCancel={cancel} />
        } />
      </Routes>
    </Container>
  )
}
