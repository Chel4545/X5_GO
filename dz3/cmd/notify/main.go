package main

import (
	"errors"
	"fmt"
	"lesson5/student/practice"
)

func printMessages(messages []string) {
	for _, message := range messages {
		fmt.Printf("%s ", message)
		fmt.Println()
	}
}

func main() {
	var email = &practice.EmailNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "email"},
		Address:      "eximple@gmail.com",
	}

	var push = &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "push"},
		DeviceID:     "1234",
	}

	var notifiers = []practice.Notifier{email, push}

	fmt.Println("ok:")

	var err = practice.Dispatch(notifiers, "order ready")

	if err == nil {
		printMessages(email.Messages)
		printMessages(push.Messages)
	}

	fmt.Println()
	fmt.Println("empt msg:")

	err = practice.Dispatch(notifiers, "")
	if errors.Is(err, practice.ErrEmptyMessage) {
		fmt.Println(err)
		printMessages(email.Messages)
		printMessages(push.Messages)
	}

	fmt.Println()
	fmt.Println("push without device id:")

	notifiers[1] = &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "push"},
	}

	err = practice.Dispatch(notifiers, "chanalname")
	var sendError *practice.SendError
	if errors.As(err, &sendError) {
		fmt.Println(err)
	}

	var brokenAndValidPushSlice = []practice.Notifier{&practice.BrokenNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "broken"},
	}, push}

	fmt.Println()
	fmt.Println("brokenNotifier and fix push:")

	var errs = practice.SafeDispatch(brokenAndValidPushSlice, "otherchanal")
	if errs != nil {
		for _, e := range errs {
			fmt.Println(e)
		}
	}
	printMessages(push.Messages)
}
