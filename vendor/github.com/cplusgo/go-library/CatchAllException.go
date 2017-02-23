package go_library

/*
模仿Java中的
	try {
		logicRun()
	} catch (exp) {
		processExp()
	}
可以有效阻止程序被异常终止，并尝试恢复正常状态
*/

type TryCatchWorker interface {
	Try()
	Catch(err interface{})
}

func Run(worker TryCatchWorker) {
	defer func() {
		if err := recover(); err != nil {
			worker.Catch(err)
		}
	}()
	worker.Try()
}
