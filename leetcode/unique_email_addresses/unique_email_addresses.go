// https://leetcode.com/problems/unique-email-addresses/
package unique_email_addresses

func numUniqueEmails(emails []string) int {
	uniqEmails := map[string]struct{}{}
	for i := 0; i < len(emails); i++ {
		email := getRealEmail(emails[i])
		uniqEmails[email] = struct{}{}
	}
	return len(uniqEmails)
}

func getRealEmail(email string) string {
	r := make([]byte, len(email))
	j := 0
	domain := false
	skip := false
	for i := 0; i < len(email); i++ {
		ch := email[i]
		if ch == '@' {
			domain = true
			skip = false
		}
		if !domain {
			if skip || ch == '.' {
				continue
			}
			if ch == '+' {
				skip = true
				continue
			}
		}

		r[j] = ch
		j++
	}
	return string(r[0:j])
}
