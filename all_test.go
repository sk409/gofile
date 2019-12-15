package gofile

import (
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	directoryPermission = 0755
)

var _ = Describe("gofile test", func() {

	var (
		err                      error
		cwd                      string
		pathTestDataExist        string
		pathTestDataCopy         string
		pathTestDataCopySrcCase0 string
		pathTestDataCopyDstCase0 string
	)

	BeforeSuite(func() {
		cwd, err = os.Getwd()
		if err != nil {
			panic(err)
		}
		pathTestDataExist = filepath.Join(cwd, "testdata.exist")
		pathTestDataCopy = filepath.Join(cwd, "testdata.copy")
		pathTestDataCopySrcCase0 = filepath.Join(pathTestDataCopy, "src", "case0")
		pathTestDataCopyDstCase0 = filepath.Join(pathTestDataCopy, "dst", "case0")
		os.MkdirAll(pathTestDataExist, directoryPermission)
		os.MkdirAll(pathTestDataCopySrcCase0, directoryPermission)
		os.MkdirAll(pathTestDataCopyDstCase0, directoryPermission)
		os.Create(filepath.Join(pathTestDataExist, "somedirectory"))
		os.Create(filepath.Join(pathTestDataExist, "somefile"))
		os.Create(filepath.Join(pathTestDataCopySrcCase0, "test1"))
		os.Create(filepath.Join(pathTestDataCopySrcCase0, "test2"))
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
			WalkFileTree(srcDirectoryPath, func(current string, fileInfo os.FileInfo) {
				dstFilePath := strings.Replace(
					filepath.Join(current, fileInfo.Name()),
					srcDirectoryPath,
					dstDirectoryPath,
					1,
				)
				isExist := IsExist(dstFilePath)
				Expect(isExist).To(BeTrue())
			})
		})
	})
})
