package token

import "github.com/google/uuid"

func HashPayloadID(id uuid.UUID, username string) uuid.UUID {
	return uuid.NewSHA1(id, []byte(username))
}

func CompareHashPayloadID(hashed_id, compare_id uuid.UUID, compare_username string) bool {
	compare_id = HashPayloadID(compare_id, compare_username)
	return hashed_id == compare_id
}
