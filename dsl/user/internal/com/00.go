package com

func Init() error {
	err := initMySql()
	if err != nil {
		return err
	}

	err = initRedis()
	if err != nil {
		return err
	}

	err = initElastic()
	if err != nil {
		return err
	}

	return nil
}
