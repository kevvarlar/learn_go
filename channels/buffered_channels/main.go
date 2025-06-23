package main

func addEmailsToQueue(emails []string) chan string {
	emailCh := make(chan string, len(emails))
	for _, email := range emails {
		emailCh <- email
	}
	return emailCh
}
