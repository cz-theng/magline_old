namespace MagNode {

	public enum ErrNO {
		Success,
	}

	public interface INode {
		ErrNO Init();

		ErrNO DeInit();

		ErrNO Connect(string URL, ulong timeout);

		ErrNO Reconnect(ulong timeout);

		ErrNO Send(byte[] data,uint length,ulong timeout);

		ErrNO Recv(ref byte[] buf,uint length,ulong timeout);

		ErrNO Close();

	}
}
