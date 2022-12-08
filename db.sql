CREATE TABLE stun (
  id uuid PRIMARY KEY,
  answer jsonb,
  answer_ice jsonb,
  offer jsonb,
  offer_ice jsonb,
  created_at date
);
