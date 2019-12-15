package gofile

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("gofile test", func() {

	var (
		err                            error
		cwd                            string
		pathTestDataExist              string
		pathTestDataCopy               string
		pathTestDataCopySrcCase0       string
		pathTestDataCopySrcCase0Case00 string
		pathTestDataCopyDstCase0       string
	)

	BeforeSuite(func() {
		mkfileWithText := func(path, text string) error {
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			defer file.Close()
			file.Write([]byte(text))
			return nil
		}
		cwd, err = os.Getwd()
		if err != nil {
			panic(err)
		}
		pathTestDataExist = filepath.Join(cwd, "testdata.exist")
		pathTestDataCopy = filepath.Join(cwd, "testdata.copy")
		pathTestDataCopySrcCase0 = filepath.Join(pathTestDataCopy, "src", "case0")
		pathTestDataCopySrcCase0Case00 = filepath.Join(pathTestDataCopySrcCase0, "case0-0")
		pathTestDataCopyDstCase0 = filepath.Join(pathTestDataCopy, "dst", "case0")
		os.MkdirAll(pathTestDataExist, directoryPermission)
		os.MkdirAll(pathTestDataCopySrcCase0Case00, directoryPermission)
		os.MkdirAll(pathTestDataCopyDstCase0, directoryPermission)
		os.Mkdir(filepath.Join(pathTestDataExist, "somedirectory"), directoryPermission)
		if err = mkfileWithText(filepath.Join(pathTestDataExist, "somefile"), "somefile"); err != nil {
			panic(err)
		}
		if err = mkfileWithText(filepath.Join(pathTestDataCopySrcCase0, "test1"), "test1"); err != nil {
			panic(err)
		}
		if err = mkfileWithText(filepath.Join(pathTestDataCopySrcCase0, "test2"), "test2"); err != nil {
			panic(err)
		}
		if err = mkfileWithText(filepath.Join(pathTestDataCopySrcCase0Case00, "test1"), "test1"); err != nil {
			panic(err)
		}
		if err = mkfileWithText(filepath.Join(pathTestDataCopySrcCase0Case00, "test2"), "test2"); err != nil {
			panic(err)
		}
	})

	AfterSuite(func() {
		os.RemoveAll(pathTestDataExist)
		os.RemoveAll(pathTestDataCopy)
	})

	Context("check for existence", func() {
		It("should be determined to exist", func() {
			isExist := IsExist(filepath.Join(pathTestDataExist, "somedirectory"))
			Expect(isExist).To(BeTrue())
		})
		It("should be determined not to exist", func() {
			IsExist := IsExist(filepath.Join(pathTestDataExist, "isnotexist"))
			Expect(IsExist).To(BeFalse())
		})
	})

	Context("copy directory to directory", func() {
		It("should not raise error when copy testdata.copy/src/case0 to testdata.copy/dst/case0", func() {
			err = Copy(
				pathTestDataCopySrcCase0,
				pathTestDataCopyDstCase0,
			)
			Expect(err).To(BeNil())
		})
		It("should exist copied files in testdata.copy/dst/case0", func() {
			srcDirectoryPath := pathTestDataCopySrcCase0
			dstDirectoryPath := pathTestDataCopyDstCase0
			// TODO: WalkFileTreeのテスト
			WalkFileTree(srcDirectoryPath, func(current string, fileInfo os.FileInfo) {
				srcFilePath := filepath.Join(current, fileInfo.Name())
				dstFilePath := strings.Replace(
					filepath.Join(current, fileInfo.Name()),
					srcDirectoryPath,
					dstDirectoryPath,
					1,
				)
				srcFile, err := os.Open(srcFilePath)
				Expect(err).To(BeNil())
				srcFileBytes, err := ioutil.ReadAll(srcFile)
				Expect(err).To(BeNil())
				dstFile, err := os.Open(dstFilePath)
				Expect(err).To(BeNil())
				dstFileBytes, err := ioutil.ReadAll(dstFile)
				Expect(err).To(BeNil())
				Expect(srcFileBytes).To(Equal(dstFileBytes))
				isExist := IsExist(dstFilePath)
				Expect(isExist).To(BeTrue())
			})
		})
	})
})
