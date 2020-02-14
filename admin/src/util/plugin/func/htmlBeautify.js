/**
 * 给html添加值定样式 常用于rich text的子节点
 * @param {*} html 需要处理的html文本
 * @param {*} rule {
			img: {
				"max-width": "100%"
			}
		}	增加的样式规则 {元素名称:{css名称:属性}}
 */
const htmlBeautify = (html, rule = {}) => {
    rule = Object.assign(rule, {
        img: {
            "max-width": "100%",
            "display": "block"
        }
    })

    for (const key in rule) {
        if (rule.hasOwnProperty(key)) {
            const element = rule[key];
            let style = ';'
            for (const k in element) {
                if (element.hasOwnProperty(k)) {
                    const v = element[k];
                    style += `${k}:${v};`
                }
            }

            // let reg = new RegExp(`<${key}>`, 'g')
            var reg = new RegExp(`(i?)(\<${key})(?!(.*?style=[\'\"](.*)[\'\"])[^\>]+\>)`, "gmi");
            let reg1 = new RegExp(`<${key}(.*?)style=[\'\"](.*?)[\'\"]`, 'gmi')


            html = html.replace(reg, `$2 style="" $3`).replace(reg1, `<${key} $1 style="$2${style}"`)
        }
    }
    // br在uniapp中显示异常
    html = html.replace(/\<\/br\>/g, '<br>')
    // section在uni.richtext中不识别
    html = html.replace(/section/g, 'div')
    
    html = html.replace(/(width=[\'\"].*?[\'\"])/g, '')
    html = html.replace(/(height=[\'\"].*?[\'\"])/g, '')
    // console.log(html)
    return html
}

export default htmlBeautify