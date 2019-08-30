pub fn explode<'a>(data:&'a str,delemiter:u8)->Vec<&'a str>{
	let mut i = 0;
	let mut begin = 0;
	let mut result = Vec::new();
	for c in data.bytes(){
		if c == delemiter || c == b' '{
			if i - begin >= 1{
				result.push(&data[begin..i]);
			}
			begin = i + 1;
		}
		i = i + 1;
	}
	if i - begin >= 1{
		result.push(&data[begin..i]);
	}
	return result;
}

pub fn implode(data:&Vec<String>,delemiter:&str)->String{
	let mut result = "".to_string();
	for single in data.iter(){
		if result.len() != 0{
			result = result + delemiter;
		}
		result = result + single;
	}
	return result;
}

//表示为单元测试，不编译进release版本
#[cfg(test)]
mod tests {
	use super::explode;
	use super::implode;

    #[test]
    fn explode_test() {
    	let vec1 = explode("  123 , abc,  wer ",b',');
    	assert_eq!(vec1, vec!["123","abc","wer"],"should be equal");
    }

    #[test]
    fn implode_test(){
    	let vec1 = vec!["123","cd","h"].into_iter().map(|data|{data.to_string()}).collect::<Vec<String>>();
    	let str1 = implode(&vec1,"_");
    	assert_eq!(str1,"123_cd_h");
    }
}
