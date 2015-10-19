namespace MagNode {

	internal class Node : INode{
		public ErrNO Init() {
			return ErrNO.Success;
		}
		
		public ErrNO DeInit() {
			return ErrNO.Success;
		}
		
		public ErrNO Connect(string URL, ulong timeout) {
			return ErrNO.Success;
		}
		
		public ErrNO Reconnect(ulong timeout) {
			return ErrNO.Success;
		}
		
		public ErrNO Send(byte[] data,uint length,ulong timeout) {
			return ErrNO.Success;
		}
		
		public ErrNO Recv(ref byte[] buf,uint length,ulong timeout) {
			return ErrNO.Success;
		}
		
		public ErrNO Close() {
			return ErrNO.Success;
		}

	}
}
