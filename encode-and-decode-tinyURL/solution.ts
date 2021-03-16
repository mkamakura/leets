const hashmap = new Map()
let i = 0

/**
 * Encodes a URL to a shortened URL.
 */
function encode(longUrl: string): string {
  i++
  hashmap.set(i, longUrl)
  return 'http://tinyurl.com/' + i
}

/**
 * Decodes a shortened URL to its original URL.
 */
function decode(shortUrl: string): string {
  return hashmap.get(+shortUrl.replace('http://tinyurl.com/', ''))
}

/**
 * Your functions will be called as such:
 * decode(encode(strs));
 */
