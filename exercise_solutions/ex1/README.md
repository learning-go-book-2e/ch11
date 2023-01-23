# Exercise 1

## Question
Go to https://www.ohchr.org/en/human-rights/universal-declaration/translations/english and copy the text of the Universal Declaration of Human Rights into a text file called english_rights.txt. Click the "Other Languages" link and copy the text of the document in a few additional languages into files named LANGUAGE_rights.txt. Create a program that embeds these files into package-level variables. Your program should take in one command-line parameter, the name of a language. It should then print out the UDHR in that language.

## Solution
Each of the embedded files needs to be embedded into its own package-level variable. This variable doesn't need to be exported,
but it does have to be of type `string` or `[]byte`. The `os.Args` variable provides access to the command-line. A `switch` 
statement allows us to choose which language to display.

