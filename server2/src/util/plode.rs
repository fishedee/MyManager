pub fn explode<'a>(data:&'a str,delimiter:byte)->Vec<&'a str>{
	let i = 0;
	let begin = 0;
	let result = Vec::new();
	for c in data.bytes(){
		if c == delimiter{
			if i - begin >= 1 {
				result.push(data[begin..i]);
				begin = i+1;
			}
		}
		i = i +1;
	}
}