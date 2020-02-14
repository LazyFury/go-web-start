const getDate = (time) => {
    let date = new Date(time)
    let year = date.getFullYear()
    let month = fix0(date.getMonth() + 1)
    let day = fix0(date.getDate())
    let h = fix0(date.getHours())
    let minutes = fix0(date.getMinutes())
    let seconds = fix0(date.getSeconds())
    return {
        year,
        month,
        day,
        h,
        minutes,
        seconds
    }
}

/**
 * 时间格式化
 * @param {*} time  10位时间戳
 * @param {*} str   制定到时间格式 y-m-d H:i:s 不区分大小写
 */
const time_format = (time, str = 'y-m-d H:i:S') => {
    let obj = getDate(time * 1)
    if (!obj) {
        return null;
    }
    let result = str.replace(/([yY])/, `${obj.year}`)
        .replace(/([mM])/, `${obj.month}`)
        .replace(/([dD])/, `${obj.day}`)
        .replace(/([hH])/, `${obj.h}`)
        .replace(/([iI])/, `${obj.minutes}`)
        .replace(/([sS])/, `${obj.seconds}`)
    return result
}

// 十以内数字补零
const fix0 = (num) => {
    return num < 10 ? String('0' + num) : String(num)
}


export default time_format