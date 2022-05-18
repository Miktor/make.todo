ACT = act -s GITHUB_TOKEN=${GITHUB_TOKEN}
WORKFLOWS_DIR = .github/workflows

ACT_PR = ${ACT} pull_request --no-skip-checkout 
ci-all:
	${ACT_PR}

WORKFLOWS=$(patsubst ${WORKFLOWS_DIR}/%.yml,%,$(shell ls -f ${WORKFLOWS_DIR}/*.yml))

$(foreach wf,$(WORKFLOWS),ci-$(wf)): ci-%:
	$(ACT_PR) -W ${WORKFLOWS_DIR}/$*.yml
