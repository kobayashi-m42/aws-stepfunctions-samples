.PHONY: push-sync
push-sync:
	cd sync && make push

.PHONY: push-wait
push-wait:
	cd waitForTaskToken && make push
