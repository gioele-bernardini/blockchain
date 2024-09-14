pub struct Block {
  // Time the block was created
  timestamp: u128,
  transactions: String,
  prev_block_hash: String,
  // Hash of the *current* block
  hash: String,
  height: usize,
  // Needeed for mining
  nonce: i32,
}

pub struct Blockchain {
  blocks: Vec<Block>,
}

impl Block {
  pub fn new_block(data: String, prev_block_hash: String, height: usize) -> Result<Block> {}
}
