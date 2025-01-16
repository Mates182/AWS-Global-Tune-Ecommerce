using Microsoft.Data.SqlClient;
using System.Data;
using System.Runtime.Serialization;
using System.ServiceModel;

namespace CreateInvoice.BusinessLogic
{
    [DataContract]
    public class Invoice
    {
        [DataMember]
        public DateTime BillingDate { get; set; }

        [DataMember]
        public string PaymentMethod { get; set; } = string.Empty;

        [DataMember]
        public decimal Amount { get; set; }

        [DataMember]
        public decimal ShippingFee { get; set; }

        [DataMember]
        public int ClientId { get; set; }

        [DataMember]
        public decimal Tax { get; set; }
    }


    [ServiceContract]
    public interface ISoapService
    {
        [OperationContract]
        int CreateInvoice(Invoice invoice);
    }

    public class SoapService : ISoapService
    {
        private readonly string _connectionString = Environment.GetEnvironmentVariable("connection");

        public int CreateInvoice(Invoice invoice)
        {
            if (invoice == null)
            {
                throw new ArgumentNullException(nameof(invoice), "Invoice cannot be null.");
            }
            AppContext.SetSwitch("Switch.Microsoft.Data.SqlClient.DisableSqlServerPerformanceCounters", true);
            Console.WriteLine(_connectionString);

            using var connection = new SqlConnection(_connectionString);
            using var command = new SqlCommand("CreateInvoice", connection)
            {
                CommandType = CommandType.StoredProcedure
            };

            // Adding parameters
            command.Parameters.AddWithValue("@billing_date", invoice.BillingDate);
            command.Parameters.AddWithValue("@payment_method", invoice.PaymentMethod);
            command.Parameters.AddWithValue("@amount", invoice.Amount);
            command.Parameters.AddWithValue("@shipping_fee", invoice.ShippingFee);
            command.Parameters.AddWithValue("@client_id", invoice.ClientId);
            command.Parameters.AddWithValue("@tax", invoice.Tax);

            connection.Open();
            var result = command.ExecuteScalar();
            connection.Close();
            return Convert.ToInt32(result); // Return the new invoice ID
        }

    }


}