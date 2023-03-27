import { Button, Container, Heading, Text } from "@chakra-ui/react";
import { Link } from "react-router-dom";

export default function OnboardIntroduction() {
  return (
    <Container>
      <Heading size="small" my="1em">
        Welcome to Octopus Jr!
      </Heading>
      <Text mb="1em">
        I'm excited to have you and your family join our community. I created this family app as a safe and fun place where my child and other children can interact with friends and family online. As a parent, I understand firsthand the importance of providing a secure environment for kids to engage in age-appropriate activities and build meaningful connections with loved ones. That's why I designed Octopus Jr. to prioritize your child's safety and security. I hope you enjoy exploring all that our family app has to offer and watching your child thrive in this positive online community.
      </Text>
      <Text mb="1em">
        I believe that building a positive online community starts with involvement from parents. As a fellow parent, I encourage you to be an active participant in the community, engaging with other families and providing feedback to help us continually improve the platform. You play a crucial role in ensuring that your child has a safe and enjoyable experience on our family app. Together, we can create a supportive and engaging community for families.
      </Text>
      <Text mb="1em">
        If you have any questions or concerns, don't hesitate to reach out.
      </Text>
      <Text>
        Ryan Yeske
      </Text>
      <address>octopusjr@ryanyeske.com</address>

      <Link to="../family">
        <Button mt="1em" mb="32" w="full" colorScheme="blue">Get Started</Button>
      </Link>
    </Container>
  );
}
