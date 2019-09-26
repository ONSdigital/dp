Contributing to Digital Publishing repos
========================================

### Development work
* We use [git-flow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) - create a feature or fix branch from develop, e.g. feature/new-feature, fix/new-fix. Alternatively, create a hotfix branch from master e.g. hot-fix/new-hotfix.
* Pull requests must use the pull request template found in each repo. Each section must contain the following:
  * 'What' a succinct, clear summary of what the user need is driving this feature change,
  * 'How to review' a step by step guide or set of guides detailing how a reviewer can test and verify the changes made
  * 'Who can review' a listing of the most appropriate people to review the pull request, e.g. if on a frontend application a member of the frontend team should review, likewise for backend and platform. If the changes are critical or likely to have severe side-effects the lead develop should review.  
* Ensure your branch contains logical atomic commits before sending a pull request - follow the [alphagov Git styleguide](https://github.com/alphagov/styleguides/blob/master/git.md)
* You may rebase your branch after feedback if it's to include relevant updates from the develop branch. We prefer a rebase here to a merge commit as we
prefer a clean and straight history on the develop branch with discrete merge commits for features
* It is advised to squash commits before pushing to a remote repository 
-----
### Releasing
* For instructions on how to release work read the [Releases guidelines](https://github.com/ONSdigital/dp/blob/master/RELEASES.md)
