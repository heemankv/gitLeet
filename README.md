
**Using the Feynman technique**

**Feynman Technique**
- Choose a topic 
	- Get something you want to learn
- Do initial research 
	- Define the scope you want to work with
- Explain it to a 12 year old
	- Document it in a paragraph as to what you want to do
- Reflect, Refine and Simplify 
	- build something fast and simple completely
- Organise and Review
	- Now find someone to review it and help you make it better 

**Choose a Topic:** 
	Get all the codes present on leetcode to github in a date based commit fashion.

**Do initial research** 
	1) **Leetcode Graphql** (https://leetcode.com/graphql/) will be needed to fetch the list of codes and code snippets.
	2) .md + .cpp file management will be needed to club all the data together in a single file and then push it. 
	3) Github git bash commands will be needed to commit code based on date received from Leetcode.

**Assumptions**
It will be a single time run service.
All the codes are written in cpp.
User is logged in to git on machine's cli 
User provides with Leetcode CSRF token


**New Tech played with**
	Extensive file management 
	Graphql api
	Github commit flow 

**Configs needed** 
Github repo r/w access

**Order looks like so:**
	Upon request
		Fetch of user's solved question list from Leetcode 
		For Each solved question we need:
			Question Id
			Solution 
			Question 
			Solution date and time
		This acquired data will be committed to Github using git bash based on  predefined dated. 

**Result:**
Github should have a repo which will have all the solutions stored as folders having  question_id.cpp and question_id.md
	question_id.cpp consists of code.
	question_id.md consists of meta data around it.


**Explaining it to a 12 year old**
**I want a single place which has all the codes written by me stored**.
I want these codes from Leetcode to help increase the language spread on my Github.


Question ? 
Which Language 
Let's take a challenge and do this in **goLang** !


**Updates and Findings:** 
1) Initially idk why I thought that it would be a rest api,  actually it will just be a go script, maybe having some bash with it, nothing else
2) "lang": 0 -> for a graphql query to leetcode cpp is lang : 0
3) for the user's solution query we are already getting timestamp of submission, this can help us is making timestamp for user commit on Github
4) Will have to run **prettier** before committing the codes.
5) Issue #1, need a repo created before the oldest time-morphed commit. because commit tree will be started from there and that is not in our hands to fake. -> https://github.com/heemankv/LeetCode-Solutions.git **SORTED**
6) Issue #2, SORT QUESTIONS BY DATE solved, so that commit graph can be made in a chronological manner !
7) Issue #3, How to commit
	1) Either assume that the user is connected to bitBash , and run bash scripts from terminal ->  https://stackoverflow.com/questions/25834277/executing-a-bash-script-from-golang
	2) Or do go-git auth process  -> https://stackoverflow.com/questions/61677396/pushing-to-a-remote-with-basic-authentication
	3) Need to do a "chmod +x pathToScript" for all the 
8) Did not account for *html* coming from graphql coming from backend assumed it to be a simple string, now might have to make a html to markdown parser or use an online tool. 
9) Do have to ensure to run prettier after each folder creation / before each commit !

**Queries :** 
Managed inside POSTMAN,
ensure to make is publicly available after removing personal info like CSRF token and so


Update: 31st July 2023
Although my code is ready till it's last 95% I am unable to understand why it's failing for this :
xD this was a simple problem wherein my solution on github was actually null and my type assertion was failing for the same, handled using assertiong validation , could have used proper error handling
