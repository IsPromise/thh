package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	Articles2 "thh/app/models/bbs/Articles"
	"thh/bundles/config"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "hexo:tool",
		Short: "",
		Run:   runHexotool,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

type Blog struct {
	Title   string
	Content string
}

func runHexotool(_ *cobra.Command, _ []string) {

	basePath := config.GetString("HEXO_POSTS_PATH", "")
	if len(basePath) == 0 {
		fmt.Println("请填写有效目标路径")
		return
	}
	blogs, err := traverse(basePath)
	if err != nil {
		fmt.Println("Error traversing directory:", err)
		return
	}
	for _, data := range blogs {
		art := Articles2.Articles{UserId: 1, Content: data.Content, Title: data.Title}
		Articles2.Save(&art)
	}
	fmt.Println(len(blogs))
}

func traverse(path string) ([]Blog, error) {
	blogs := make([]Blog, 0)

	files, err := os.ReadDir(path)
	if err != nil {
		return blogs, err
	}

	for _, file := range files {
		if file.IsDir() {
			subPath := filepath.Join(path, file.Name())
			subBlogs, err := traverse(subPath)
			if err != nil {
				return blogs, err
			}
			blogs = append(blogs, subBlogs...)
		} else {
			if strings.HasSuffix(file.Name(), ".md") {
				filePath := filepath.Join(path, file.Name())
				data, err := os.ReadFile(filePath)
				if err != nil {
					fmt.Println("Error reading file:", err)
					continue
				}

				title := ""
				content := ""

				// Parse front-matter and content
				scanner := bufio.NewScanner(strings.NewReader(string(data)))
				isFrontMatter := true
				for scanner.Scan() {
					line := scanner.Text()
					if line == "---" {
						if isFrontMatter {
							isFrontMatter = false
						} else {
							break
						}
					}
					if strings.HasPrefix(line, "title:") {
						title = strings.TrimSpace(strings.TrimPrefix(line, "title:"))
					}
				}
				for scanner.Scan() {
					content += scanner.Text() + "\n"
				}

				blogs = append(blogs, Blog{
					Title:   title,
					Content: content,
				})
			}
		}
	}

	return blogs, nil
}
