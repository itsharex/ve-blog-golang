/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package model

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ve-weiyi/ve-blog-golang/quickstart/cmd/model/helper"
)

// migrateCmd represents the migrate command
type ModelDSNCmd struct {
	CMD *cobra.Command
	modelConfig
}

func NewModelDSNCmd() *ModelDSNCmd {
	rootCmd := &ModelDSNCmd{}
	rootCmd.CMD = &cobra.Command{
		Use: "dsn",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.RunCommand(cmd, args)
		},
	}

	rootCmd.init()
	return rootCmd
}

func (s *ModelDSNCmd) init() {
	dsn := "root:mysql7914@(127.0.0.1:3306)/blog-veweiyi?charset=utf8mb4&parseTime=True&loc=Local"
	s.CMD.PersistentFlags().StringVarP(&s.SqlFile, "source", "s", dsn, "数据库地址")
	s.CMD.PersistentFlags().StringVarP(&s.TplFile, "tpl-file", "t", "entity.tpl", "模板文件")
	s.CMD.PersistentFlags().StringVarP(&s.OutPath, "out-path", "o", "./", "输出路径")
	s.CMD.PersistentFlags().StringVarP(&s.NameAs, "name-as", "n", "%s.go", "输出名称")
}

func (s *ModelDSNCmd) RunCommand(cmd *cobra.Command, args []string) {
	log.Println("run model ddl")
	log.Println("sql-file:", s.SqlFile)
	log.Println("tpl-file:", s.TplFile)
	log.Println("out-path:", s.OutPath)
	log.Println("name-as:", s.NameAs)

	var tables []*helper.Table
	var err error

	tables, err = helper.ParseTableFromDsn(s.SqlFile)
	if err != nil {
		panic(err)
	}

	err = generateModel(tables, s.modelConfig)
	if err != nil {
		panic(err)
	}
}